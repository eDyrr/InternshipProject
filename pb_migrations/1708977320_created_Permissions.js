/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "0cdu8hqot90pfre",
    "created": "2024-02-26 19:55:20.318Z",
    "updated": "2024-02-26 19:55:20.318Z",
    "name": "Permissions",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "xl0lnuvg",
        "name": "permission",
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
  const collection = dao.findCollectionByNameOrId("0cdu8hqot90pfre");

  return dao.deleteCollection(collection);
})
