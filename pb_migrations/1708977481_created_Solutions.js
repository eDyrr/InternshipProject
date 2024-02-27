/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "7ycopou9s9ccwbs",
    "created": "2024-02-26 19:58:01.375Z",
    "updated": "2024-02-26 19:58:01.375Z",
    "name": "Solutions",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ohh1j0jg",
        "name": "Solution",
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
  const collection = dao.findCollectionByNameOrId("7ycopou9s9ccwbs");

  return dao.deleteCollection(collection);
})
