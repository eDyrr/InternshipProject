/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7ycopou9s9ccwbs")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dmmajb9t",
    "name": "ticket",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "popy1zum3qa1yap",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7ycopou9s9ccwbs")

  // remove
  collection.schema.removeField("dmmajb9t")

  return dao.saveCollection(collection)
})
