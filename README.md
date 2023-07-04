# Youtube API Server
## Overview 
This launches a golang server that can fetch videos using youtube-api and stores the results in a database(mysql) .

## Examples
### Searching Video 
` http://localhost:3000/api/search_videos?query=<YOUR_QUERY>&count=<ITEM_COUNT> `
<br>
<br>
![Alt text](assets/search.png)

### Fetching Videos
` http://localhost:3000/api/_videos?query=<YOUR_QUERY>&count=<ITEM_COUNT> `
<br>
<br>
![Alt text](assets/get.png)
