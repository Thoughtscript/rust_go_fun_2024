db = db.getSiblingDB('testdatabase')

db.createUser({
  user: 'testuser',
  pwd: 'testpass',
  roles: [
    { role: 'readWrite', db: 'testdatabase' }
  ]
})

db.createCollection('examples')