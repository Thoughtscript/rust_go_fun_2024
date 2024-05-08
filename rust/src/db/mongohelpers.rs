use mongodb::{Client, Collection};

use crate::domain::example::{Example};

pub struct ExampleMongoHelper {
    pub collection: Collection<Example>
}

// define methods - unlike go, no receiver functions
// like C++ with :: scope resolution operator 
impl ExampleMongoHelper {

    // initialize once with a single client
    // https://github.com/Mr-Malomz/actix-mongo-api/blob/f40ee0c761c1d8d54941ee103652d9487a6dcaae/src/repository/mongodb_repo.rs#L26-L30
    pub async fn init() -> Self {
        let client = Client::with_uri_str("mongodb://testuser:testpass@mongodb:27017/testdatabase")
        .await
        .expect("error connecting to database");
        
        let db = client.database("testdatabase");
        let collection: Collection<Example> = db.collection("examples");
        ExampleMongoHelper { collection } // the absence of a semi-colon returns 
    }
}
