/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "n0z0i0ac",
    "name": "owner",
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
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  // remove
  collection.schema.removeField("n0z0i0ac")

  return dao.saveCollection(collection)
})
