/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "d5hrmbhdk8odtpl",
    "created": "2024-11-14 05:11:49.661Z",
    "updated": "2024-11-14 05:11:49.661Z",
    "name": "list",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ugspsuxd",
        "name": "email",
        "type": "email",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "exceptDomains": null,
          "onlyDomains": null
        }
      },
      {
        "system": false,
        "id": "ohjysreq",
        "name": "name",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("d5hrmbhdk8odtpl");

  return dao.deleteCollection(collection);
})
