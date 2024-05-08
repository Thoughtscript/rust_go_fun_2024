use serde::{Serialize, Deserialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Example {
    pub id: String,
    pub name: String,
    pub val: String,
}