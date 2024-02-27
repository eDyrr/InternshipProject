/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "o7ayag9b",
    "name": "solution",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "7ycopou9s9ccwbs",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  // remove
  collection.schema.removeField("o7ayag9b")

  return dao.saveCollection(collection)
})
