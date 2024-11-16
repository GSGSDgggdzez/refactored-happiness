/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "az50yaakht3bchr",
    "created": "2024-11-14 08:51:19.926Z",
    "updated": "2024-11-14 08:51:19.926Z",
    "name": "waitlist",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "bkyjuw4d",
        "name": "email",
        "type": "email",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "exceptDomains": [],
          "onlyDomains": []
        }
      },
      {
        "system": false,
        "id": "udiuhfw6",
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
    "listRule": "",
    "viewRule": "",
    "createRule": "",
    "updateRule": "",
    "deleteRule": "",
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("az50yaakht3bchr");

  return dao.deleteCollection(collection);
})
