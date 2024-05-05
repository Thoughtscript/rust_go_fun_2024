use actix_web::{web, App, HttpServer};

mod api; // Local files in same dir are imported this way

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(api::hello)
            .service(api::echo)
            .route("/hey", web::get().to(api::manual_hello))
    })
    .bind(("0.0.0.0", 8000))?
    .run()
    .await
}