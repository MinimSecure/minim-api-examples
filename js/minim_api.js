import axios from "axios";

axios.defaults.headers = { 'User-Agent': 'Minim API Examples' };

export default {

  api_endpoint: process.env.MINIM_URL || 'https://my.minim.co',

  async init() {
    this.api_secret = process.env.API_SECRET
    this.api_app_id = process.env.APPLICATION_ID
    await this.getAccessToken();
  },

  /**
   * Fetches ALL the paginated IDs from a given api/v1 INDEX endpoint
   *
   * @param {*} url - INDEX endpoint for a controller - ex: api/v1/lans/{lan_id}/devices
   */
  async fetch_all_ids(url, opts={}) {
    let aggregate_data = [];
    let offset = 0;
    let finished = false;

    while(!finished) {
      Object.assign(opts, {offset: offset})
      const res = await this.get(url, opts);
      const ids = res.data.map(({ id }) => id);

      aggregate_data = aggregate_data.concat(ids);

      if (aggregate_data.length >= parseInt(res.headers['x-total-count'])) {
        finished = true;
      } else {
        offset = aggregate_data.length;
      }
    }

    return aggregate_data;
  },

  /**
   * Uses the multi-get functionality to fetch all the resources from a given api/v1 INDEX endpoint
   *
   * @param {String} url - INDEX endpoint for a controller - ex: api/v1/lans/{lan_id}/devices
   * @param {Object} opts - optional arguments - see usage below
   */
  async multi_get(
    url,
    opts = {
      max_multiget_limit: 10,
      callback: null,
      ids: null,
      params: {}
    }
  ) {
    const ids = opts.ids || await this.fetch_all_ids(url); // default to fetching all the IDs for a given resource
    const max_multiget_limit = opts.max_multiget_limit || 10;

    let ids_in_chunks = [];
    let aggregate_data = [];

    for(let i = 0; i < ids.length; i += max_multiget_limit) {
      ids_in_chunks.push(ids.slice(i, i + max_multiget_limit));
    }

    for (let i = 0; i < ids_in_chunks.length; i++) {
      // Re-assign params to be an object that contains whatever was in the
      // original params object but override the current chunk of IDs
      const chunk = ids_in_chunks[i];
      const params = { ...(opts.params || {}), id: chunk.join(',') };

      const res = await this.get(url, params);

      // Pass the response to a callback if we have one
      if (opts.callback) {
        opts.callback(res);
      }

      aggregate_data = [...aggregate_data, ...res.data];
    }

    return aggregate_data;
  },

  get(url, opts) {
    return this.request('get', url, opts);
  },

  post(url, data) {
    return this.request('post', url, null, data);
  },

  put(url, data) {
    return this.request('put', url, null, data);
  },

  patch(url, data) {
    return this.request('patch', url, null, data);
  },

  delete(url, opts) {
    return this.request('delete', url, opts);
  },

  request(method, url, opts, data = null) {
    opts = opts || {};
    return axios({
      method,
      baseURL: this.api_endpoint,
      url,
      data,
      params: Object.assign(opts, {access_token: this.access_token})
    });
  },

  async getAccessToken() {
    let data = {
      client_id: this.api_app_id,
      client_secret: this.api_secret,
      grant_type: 'client_credentials'
    };

    return axios({
      method: 'post',
      url: `${this.api_endpoint}/api/oauth/token`,
      data: data,
    }).then((response) => {
      //handle success
      let data = response.data;
      this.access_token = data.access_token;

      //setTimeout(this.getAccessToken.bind(this), ( data.expires_in*1000 - 1000 ));
    }).catch((response) => {
      // TODO - handle error
      console.log(response);
      console.log('unable to get token... exiting');
      process.exit(1);
    });
  },
};
