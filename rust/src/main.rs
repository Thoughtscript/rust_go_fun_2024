use actix_web::{web, App, HttpServer};

// note how each module/directory below has a mod.rs
// use mod here to load handlers directory in main.rs
mod handlers; // precisely this syntac - mod module_name; (with a colon!)
mod db;
mod domain;
// define all the modules here
// elsewhere use crate:: notation

// https://doc.rust-lang.org/reference/attributes.html
// attribute -metadata
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    //this is app data - not persisted
    //let db = mongohelpers::get_collection().await;
    //let db_data = Data::new(db);

    HttpServer::new(|| {
        App::new()
            //this is app data - not persisted
            //.app_data(db_data.clone())
            
            //examples from actix
            .service(handlers::basicapi::hello)
            .service(handlers::basicapi::echo)
            .route("/hey", web::get().to(handlers::basicapi::manual_hello))
            
            // Mongo
            .service(handlers::mongoapi::create_example)
            .service(handlers::mongoapi::delete_example)
            .service(handlers::mongoapi::update_example)
            .service(handlers::mongoapi::get_example)
            .service(handlers::mongoapi::get_examples)
    })
    .bind(("0.0.0.0", 8000))?
    .run()
    .await
}