# Fetching LAN IDs and Updating a LAN

Many of the calls in the Minim API require IDs to specify single-target actions. For example, LAN IDs are required to make an update to particular LANs in the database. Using Unums from `example.js` as a framework, this guide covers how to create similar variables that store LANs as an array of objects and how to extract a array of LAN IDs.

We recommend commenting out lines 14-19 while experimenting with other API calls to decrease loading time.

## Storing a list of LANs in a variable

Within `examples.js`, add the following code inside the asynchronous function under line 19:

```javascript
let lans = await minim_api.multi_get(`api/v1/lans`);
console.log(`Found ${lans.length} IDs, fetching details...`);
console.log(lans);
```

A successful output:

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

Using [`map()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map), we can create an array of the IDs (or other values) to use in other API calls detailed in the Minim API documentation.

Within `example.js`, add the following code inside the asynchronous function under line 22:

```javascript
console.log(`Fetching LAN IDs...`);  
let lanIDs = lans.map(l => l.id);
console.log(lanIDs);
```

A successful output:

```javascript
Fetching LAN IDs...
[ '081c2538-36ff-4ddb-a8b6-8cf0013b03b7' ]
```

You can also customize the IDs stored in an array using functions that map the array by city, state, speed tier, and more.

## Updating a LAN by ID

Changes to LANs are made using POST/PATCH requests to the the URLs specified in the API documentation. As an example, let's look at how to update a customer name with a corresponding LAN ID.

Within `example.js`, add the following code inside the asynchronous function:

```javascript
minim_api.patch(`api/v1/lans/${lanIDs}`, {customer_name: 'My Home Network'});
console.log(lans.customer_name);
```

Note that the example array of LANs returned only one LAN with a single ID. The `${lanIDs}` template literal is used for sake of syntactical simplicity. You can replace `lanIDs` with a single ID or a custom array of IDs. 

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

