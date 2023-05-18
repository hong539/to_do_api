use actix_web::{get, post, put, delete, web, App, HttpResponse, HttpServer, Responder};
use mysql::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct User {
    id: Option<u32>,
    name: String,
}

fn establish_connection() -> Pool {
    let url = "mysql://user:password@localhost/database";
    let pool = Pool::new(url).unwrap();
    pool
}

#[get("/users")]
async fn get_users() -> impl Responder {
    let pool = establish_connection();
    let mut conn = pool.get_conn().unwrap();

    let users: Vec<User> = conn.query_map(
        "SELECT id, name FROM users",
        |(id, name)| User { id, name },
    ).unwrap();

    HttpResponse::Ok().json(users)
}

#[post("/users")]
async fn create_user(user: web::Json<User>) -> impl Responder {
    let pool = establish_connection();
    let mut conn = pool.get_conn().unwrap();

    conn.exec_drop(
        "INSERT INTO users (name) VALUES (?)",
        (user.name.clone(),),
    ).unwrap();

    HttpResponse::Ok().json(user.into_inner())
}

#[put("/users/{id}")]
async fn update_user(web::Path(id): web::Path<u32>, user: web::Json<User>) -> impl Responder {
    let pool = establish_connection();
    let mut conn = pool.get_conn().unwrap();

    conn.exec_drop(
        "UPDATE users SET name = ? WHERE id = ?",
        (user.name.clone(), id),
    ).unwrap();

    HttpResponse::Ok().json(user.into_inner())
}

#[delete("/users/{id}")]
async fn delete_user(web::Path(id): web::Path<u32>) -> impl Responder {
    let pool = establish_connection();
    let mut conn = pool.get_conn().unwrap();

    conn.exec_drop("DELETE FROM users WHERE id = ?", (id,)).unwrap();

    HttpResponse::Ok()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new()
        .service(get_users)
        .service(create_user)
        .service(update_user)
        .service(delete_user))
        .bind("127.0.0.1:8080")?
        .run()
        .await
}