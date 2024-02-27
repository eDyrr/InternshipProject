/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  collection.name = "Tickets"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("popy1zum3qa1yap")

  collection.name = "Ticket"

  return dao.saveCollection(collection)
})
