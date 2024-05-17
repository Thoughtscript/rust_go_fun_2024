use actix_web::{web, App, HttpServer};

// note how each module/directory below has a mod.rs
// use mod here to load handlers directory in main.rs
mod handlers; // precisely this syntac - mod module_name; (with a colon!)
mod db;
mod domain;
// define all the modules here
// elsewhere use crate:: notation

use db::mongohelpers::ExampleMongoHelper;
use handlers::{basicapi, mongoapi};

// https://doc.rust-lang.org/reference/attributes.html
// attribute -metadata
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // this is app data - it actually persists to Mongo (and not to the in-memory store alone!)
    // you can test this by creating entities through the Rust API then calling through the Go API
    let mdb = ExampleMongoHelper::init().await;
    let db_data = web::Data::new(mdb);

    // https://doc.rust-lang.org/std/keyword.move.html
    HttpServer::new(move || {
        App::new()
            //this is app data
            .app_data(db_data.clone())
            
            //examples from actix
            .service(basicapi::hello)
            .service(basicapi::echo)
            .route("/hey", web::get().to(basicapi::manual_hello))
            
            // Mongo
            .service(mongoapi::create_example)
            .service(mongoapi::delete_example)
            .service(mongoapi::update_example)
            .service(mongoapi::get_example)
            .service(mongoapi::get_examples)
    })
    .bind(("0.0.0.0", 8000))?
    .run()
    .await
}