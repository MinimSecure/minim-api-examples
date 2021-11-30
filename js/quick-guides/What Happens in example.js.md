# What Happens in example.js

The most notable action in `example.js` occurs in lines 14-21 where two important variables are declared:

```javascript
  let isp_id = isp_ids[0];
  console.log(`Fetching IDs for ${isp_id}...`);
  let unum_ids = await minim_api.fetch_all_ids(`api/v1/isps/${isp_id}/unums`);
  console.log(`Found ${unum_ids.length} IDs, fetching details...`);
  let unums = await minim_api.multi_get(`api/v1/isps/${isp_id}/unums`, { ids: unum_ids, })
```

Each variable stores data that can be used for further API interaction:

| Variable     | Purpose                                          | Data Type          |
| ------------ | ------------------------------------------------ | ------------------ |
| **`isp_id`** | Retrieves the unique ID tied to your database    | *String*           |
| **`unums`**  | A list of Minim-enabled routers separated by IDs | *Array of Objects* |

## Notes for using these variables alongside [Minim Docs](https://my.minim.co/api_doc)

Minim hosts a full list of executable HTTP requests, some of which require IDs to specify particular unums, LANs, people, etc. In the case of the ISP ID example, template literals are used for simplicity. 

You can provide particular IDs in your requests to target specific ISPs. For example, adding a speed tier is performed via a POST request to `my.minim.com/v1/isps/{ispId}/speed_tiers`.

`example.js` creates the `unums` variable for you. Using this framework, you can also create a variable for LAN IDs, users, devices, and other useful values. See [Fetching LAN IDs and Updating a LAN](https://github.com/MinimSecure/minim-api-examples/blob/main/js/quick-guides/Fetching%20LAN%20IDs%20and%20Updating%20a%20LAN.md) for more details and an example.
