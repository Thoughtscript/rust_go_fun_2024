use actix_web::{get, post, delete, put, HttpResponse, Responder};

use crate::domain::example::{Example};
use crate::db::mongohelpers::{get_collection};

#[get("/examples")]
pub async fn get_examples() -> impl Responder {
    let collection = get_collection();

    HttpResponse::Ok().body("Hello world!")
}

#[get("/example/one")]
pub async fn get_example() -> impl Responder {
    let collection = get_collection();
    
    HttpResponse::Ok().body("Hello world!")
}

#[delete("/example/delete")]
pub async fn delete_example() -> impl Responder {
   let collection = get_collection();
    
    HttpResponse::Ok().body("Hello world!")
}

#[post ("/example/create")]
pub async fn create_example() -> impl Responder {
    let collection = get_collection();

    HttpResponse::Ok().body("Hello world!")
}

#[put ("/example/update")]
pub async fn update_example() -> impl Responder {
    let collection = get_collection();

    HttpResponse::Ok().body("Hello world!")
}

