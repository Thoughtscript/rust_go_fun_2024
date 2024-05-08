use actix_web::{web::{Data, Query}, get, post, delete, put, HttpResponse};
use futures::stream::{TryStreamExt};
use serde::{Deserialize};
use mongodb::bson::{doc};

use crate::domain::example::Example;
use crate::db::mongohelpers::ExampleMongoHelper;

// https://actix.rs/docs/extractors/#query
#[derive(Deserialize)]
struct IdParam {
    id: String,
}

#[derive(Deserialize)]
struct ExampleParams {
    id: String,
    name: String,
    val: String,
}

#[get("/examples")]
pub async fn get_examples(db: Data<ExampleMongoHelper>) -> HttpResponse {
    let cursor = db.collection.find(None, None).await.expect("error");

    // https://doc.rust-lang.org/std/vec/struct.Vec.html
    let results : Vec<Example> = cursor.try_collect().await.unwrap_or_else(|_| vec![]);

    HttpResponse::Ok().json(results)
}

// https://stackoverflow.com/questions/54406029/how-can-i-parse-query-strings-in-actix-web
#[get("/example/one")]
pub async fn get_example(db: Data<ExampleMongoHelper>, params: Query<IdParam>) -> HttpResponse {       
    let id = &params.id;
    
    let filter = doc! {"id": id};

    let example = db.collection.find_one(filter, None).await.expect("error");

    HttpResponse::Ok().json(example)
}

#[delete("/example/delete")]
pub async fn delete_example(db: Data<ExampleMongoHelper>, params: Query<IdParam>) -> HttpResponse {
    let id = &params.id;

    let filter = doc! {"id": id};

    let example = db.collection.delete_one(filter, None).await.expect("error");

    HttpResponse::Ok().json(example)
}

#[post("/example/create")]
pub async fn create_example(db: Data<ExampleMongoHelper>, params: Query<ExampleParams>) -> HttpResponse {
    let id = &params.id;
    let name = &params.name;
    let val = &params.val;

    let example = Example {
        id: id.to_string(),
        name: name.to_string(),
        val: val.to_string(),
    };

    // TODO - verify if doesn't exist first - or upsert?
    let result = db.collection.insert_one(example, None).await.expect("error");

    HttpResponse::Ok().json(result)
}

#[put("/example/update")]
pub async fn update_example(db: Data<ExampleMongoHelper>, params: Query<ExampleParams>) -> HttpResponse {
    let id = &params.id;
    let name = &params.name;
    let val = &params.val;

    let filter = doc! {"id": id};

    let update = doc! {
        "$set":
            {
                "id": id.to_string(),
                "name": name.to_string(),
                "val": val.to_string()
            },
    };

    let example = db.collection.update_one(filter, update, None).await.expect("error");

    HttpResponse::Ok().json(example)
}

