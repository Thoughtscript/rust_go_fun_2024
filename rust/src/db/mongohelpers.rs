use mongodb::{
    options::{ClientOptions, Credential},
    Client, Collection,
};

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
        // avoids strict Rust authSource and: "ProtocolError","errmsg":"Attempt to switch database target during SASL authentication..." issues
        let mut client_options = ClientOptions::parse("mongodb://mongodb:27017")
            .await
            .expect("error parsing MongoDB options");

        let credential = Credential::builder()
            .username(Some("testuser".to_owned()))
            .password(Some("testpass".to_owned()))
            .source(Some("testdatabase".to_owned())) // authSource and DB in one setting
            .build();

        client_options.credential = Some(credential);

        let client = Client::with_options(client_options)
            .expect("error connecting to database");
        
        let db = client.database("testdatabase");
        let collection: Collection<Example> = db.collection("examples");
        ExampleMongoHelper { collection } // the absence of a semi-colon returns 
    }
}
