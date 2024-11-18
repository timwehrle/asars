use reqwest::{
    header::{HeaderMap, AUTHORIZATION},
    Client,
};
use serde::Deserialize;
use std::error::Error;

fn asana_api_url(endpoint: &str) -> String {
    format!("https://app.asana.com/api/1.0/{}", endpoint)
}

#[derive(Deserialize, Debug)]
pub struct Workspace {
    pub gid: String,
    pub name: String,
}

pub async fn fetch_workspaces(token: &str) -> Result<Vec<Workspace>, Box<dyn Error>> {
    let client = Client::new();
    let mut headers = HeaderMap::new();
    headers.insert(AUTHORIZATION, format!("Bearer {}", token).parse()?);

    let url = format!("{}", asana_api_url("workspaces"));

    let response = client
        .get(&url)
        .headers(headers)
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    let workspaces = serde_json::from_value(response["data"].clone())?;
    Ok(workspaces)
}
