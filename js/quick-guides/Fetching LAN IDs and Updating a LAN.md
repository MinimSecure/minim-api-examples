# Fetching LAN IDs and Updating a LAN

Before proceeding, please follow the instructions in [README.md](/README.md) to establish your Minim account with an App ID and Secret Key, enabling you to retrieve, update, and delete data via the Minim API.

Many of the calls in the Minim API require IDs to specify single-target actions. For example, LAN IDs are required to make an update to specified LANs in the database. Using `unums` from `example.js` as a framework, this guide demonstrates how to declare a variable that stores LANs as an array of objects and extract an array of LAN IDs.

We recommend commenting out lines 14-19 while experimenting with other API calls to decrease loading time.

## Storing a list of LANs in a variable

Within `examples.js`, add the following code inside the asynchronous function:

```javascript
let lans = await minim_api.multi_get(`api/v1/lans`);
console.log(`Found ${lans.length} IDs, fetching details...`);
console.log(lans);
```

A successful output stores LANs as an array of objects:

```javascript
Found 1 IDs, fetching details...
[
  {
    id: '081c2538-36ff-4ddb-a8b6-8cf0013b03b7',
    href: 'https://my.minim.co/api/v1/lans/081c2538-36ff-4ddb-a8b6-8cf0013b03b7',
    customer_name: 'My Home',
...
```

## Storing an array of LAN IDs

LAN IDs are used to target desired networks to receive commands such as pausing and unpausing connectivity, sending push notifications, retrieving device connection reports, and more. Using [`map()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map), we can create an array of the IDs to use in other API calls detailed in the [Minim API documentation](https://my.minim.co/api_doc).

Within `example.js`, add the following code inside the asynchronous function:

```javascript
console.log(`Fetching LAN IDs...`);  
let lanIDs = lans.map(l => l.id);
console.log(lanIDs);
```

A successful output stores the LAN IDs as an array of strings:

```javascript
Fetching LAN IDs...
[ '081c2538-36ff-4ddb-a8b6-8cf0013b03b7' ]
```

You can also customize the IDs stored in an array using functions that map the array by city, state, speed tier, and more.

## Updating a LAN by ID

Changes to LANs are made using POST/PATCH requests to the the URLs specified in the [API documentation](https://my.minim.co/api_doc). As an example, let's look at how to update a customer name with a corresponding LAN ID.

Within `example.js`, add the following code inside the asynchronous function:

```javascript
minim_api.patch(`api/v1/lans/081c2538-36ff-4ddb-a8b6-8cf0013b03b7`, {customer_name: 'My Home Network'});
console.log(lans.customer_name);
```

Note that the example array of LANs holds only one LAN with a single ID. You can replace `lanIDs` with a single ID or a custom array of IDs. 

A successful output shows the `customer_name` variable has been changed from 'My Home' to 'My Home Network':

```javascript
Found 1 IDs, fetching details...
[
  {
    id: '081c2538-36ff-4ddb-a8b6-8cf0013b03b7',
    href: 'https://my.minim.co/api/v1/lans/081c2538-36ff-4ddb-a8b6-8cf0013b03b7',
    customer_name: 'My Home Network',
...
```
