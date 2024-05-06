use mongodb::{Client, options::ClientOptions, Collection};

use crate::domain::example::{Example};

pub async fn get_collection() -> Collection<Example> {
    // https://github.com/Mr-Malomz/actix-mongo-api/blob/f40ee0c761c1d8d54941ee103652d9487a6dcaae/src/repository/mongodb_repo.rs#L26-L30
    let client = Client::with_uri_str("mongodb://testuser:testpass@mongodb:27017/testdatabase")
        .await
        .expect("error connecting to database");

    let db = client.database("testdatabase");

    let collection: Collection<Example> = db.collection("examples");
    collection// no semicolon returns
}