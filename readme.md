# roller
git remote add deploy dokku@ultilabs.xyz:roller-api

background info

 * http://stevehanov.ca/blog/index.php?id=132

----
## new campaign
POST https://{}/api/campaign

* name: name of the campaign to track
* options: options available to select from

```json
{
	"name": "MyTest",
	"options": ["Joseph","Cason","Alicia","Chris"]
}
```

----
## Roll the dice
GET https://{}/api/campaign/{campaign name}/roll

* Will give you an option based on algorithm from article

----
## Post selection
POST https://{}/api/selection/{campaign name}/{option name selected}

* Track when an option is selected
