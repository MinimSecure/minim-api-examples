import dotenv from 'dotenv';

dotenv.config();

import minim_api from './minim_api';

(async () => {
  await minim_api.init()
  let isp_ids = await minim_api.fetch_all_ids(`api/v1/isps`);
  if(!isp_ids || !isp_ids.length) {
    console.log('No ISPs were found under your user account!');
    process.exit(1);
  }
  let isp_id = isp_ids[0];
  console.log(`Fetching IDs for ${isp_id}...`);
  let unum_ids = await minim_api.fetch_all_ids(`api/v1/isps/${isp_id}/unums`);
  console.log(`Found ${unum_ids.length} IDs, fetching details...`);
  let unums = await minim_api.multi_get(`api/v1/isps/${isp_id}/unums`, { ids: unum_ids, })
  console.log(unums)
})();
