/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "popy1zum3qa1yap",
    "created": "2024-02-26 19:57:46.400Z",
    "updated": "2024-02-26 19:57:46.400Z",
    "name": "Ticket",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "9vlqnbki",
        "name": "title",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "an9w4cd2",
        "name": "description",
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
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap");

  return dao.deleteCollection(collection);
})
