use reqwest::Client;
use serde_json::Value;

pub async fn list_projects(token: &str) -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new();
    let response = client
        .get("https://app.asana.com/api/1.0/projects")
        .bearer_auth(token)
        .send()
        .await?;

    let json: Value = response.json().await?;
    println!("{:#?}", json);
    Ok(())
}