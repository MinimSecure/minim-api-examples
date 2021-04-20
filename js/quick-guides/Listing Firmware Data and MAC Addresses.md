# Listing Firmware Data and MAC Addresses

Before proceeding, please follow the instructions in [README.md](README.md) to establish your Minim account with an App ID and Secret Key, enabling you to retrieve, update, and delete data via the Minim API. [Click here](What Happens in example.js.md) for general tips and information on `example.js`.

As an example, we provide this guide on fetching MAC addresses and firmware versions from your database. Since `unums` is an array of objects containing all client routers, we can declare a custom variable to specify data and then use `map()` to extract that data. Replace line 19 (`console.log(unums)`) in `example.js` with the following code:

```javascript
let lanNetworkInfo = unums.map((a) => [
  a.lan_mac_address,
  a.firmware_version,
  a.is_firmware_update_available,
]);
console.log(lanNetworkInfo);
```

A successful output:

```javascript
> minim-api-examples@1.0.0 start
> node -r esm example.js

Fetching IDs for minim-prototypes...
Found 401 IDs, fetching details...
[
  [ 'f4:f2:6d:63:9e:76', '180312.1322.12.0', true ],
  [ '18:d6:c7:6d:17:9c', '1.2.0.105', true ],
  [ '84:16:f9:25:6f:e1', '1.2.0.431', false ],
  [ '18:d6:c7:6d:6a:a6', '1.0.0.24', true ],
  [ '18:d6:c7:8c:e2:cc', '1.2.0.149', true ],

  ...
```

## Other useful values

Objects inside `unums` contain a variety of useful values that are fully detailed in [Minim documentation](https://my.minim.co/api_doc). Here are a few examples:

| value                         | Purpose                                                  | data type |
| ----------------------------- | -------------------------------------------------------- | --------- |
| **`firmware_update_version`** | Latest version of Minim software available for this Unum | _String_  |
| **`last_online_at`**          | Last time Unum was seen on LAN                           | _String_  |
| **`firmware_updated_at`**     | Time of last recorded firmware update                    | _String_  |
| **`is_rebooting`**            | Displays if hardware is still mid-reboot                 | _Boolean_ |
