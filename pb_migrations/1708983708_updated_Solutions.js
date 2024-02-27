/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7ycopou9s9ccwbs")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "zfifwnn3",
    "name": "writer",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
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
  collection.schema.removeField("zfifwnn3")

  return dao.saveCollection(collection)
})
