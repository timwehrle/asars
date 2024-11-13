use crate::asana_api::{fetch_workspaces, Workspace};
use directories::ProjectDirs;
use keyring::Entry;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fs;
use std::io::{self, Write};

#[derive(Serialize, Deserialize)]
pub struct Config {
    pub(crate) workspace: String,
}

pub async fn handle_config_command(delete: bool) -> Result<(), Box<dyn std::error::Error>> {
    if delete {
        println!("Deleting token...");
        delete_token().map(|msg| println!("{}", msg))?;
        return Ok(());
    }

    configure().await
}

pub(crate) async fn configure() -> Result<(), Box<dyn std::error::Error>> {
    let token = prompt_for_token()?;
    save_token(&token)?;

    let workspaces = fetch_workspaces(&token).await?;
    display_workspaces(&workspaces);

    let selected_workspace = select_workspace(&workspaces)?;
    save_config(&Config {
        workspace: selected_workspace.gid.clone(),
    })?;

    println!("Default workspace '{}' saved.", selected_workspace.name);
    Ok(())
}

fn prompt_for_token() -> Result<String, Box<dyn std::error::Error>> {
    print!("Enter your Asana Personal Access Token: ");
    io::stdout().flush()?;
    let mut token = String::new();
    io::stdin().read_line(&mut token)?;
    Ok(token.trim().to_string())
}

fn save_token(token: &str) -> Result<(), Box<dyn std::error::Error>> {
    let entry = Entry::new("asars", "token")?;
    entry.set_password(token)?;
    println!("Token saved.");
    Ok(())
}

fn display_workspaces(workspaces: &[Workspace]) {
    println!("Available workspaces:");
    for (i, workspace) in workspaces.iter().enumerate() {
        println!("{}. {}", i + 1, workspace.name);
    }
}

fn select_workspace(workspaces: &[Workspace]) -> Result<&Workspace, Box<dyn std::error::Error>> {
    print!("Select the number of the default workspace: ");
    io::stdout().flush()?;
    let mut choice = String::new();
    io::stdin().read_line(&mut choice)?;
    let choice: usize = choice.trim().parse()?;

    workspaces
        .get(choice - 1)
        .ok_or_else(|| "Invalid choice.".into())
}

fn config_path() -> Result<std::path::PathBuf, Box<dyn std::error::Error>> {
    if let Some(proj_dirs) = ProjectDirs::from("com", "asars", "cli") {
        Ok(proj_dirs.config_dir().join("config.json"))
    } else {
        Err("Failed to locate configuration file directory.".into())
    }
}

fn save_config(config: &Config) -> Result<(), Box<dyn std::error::Error>> {
    let path = config_path()?;
    fs::create_dir_all(path.parent().unwrap())?;
    let config_data = serde_json::to_string_pretty(config)?;
    fs::write(path, config_data)?;
    Ok(())
}

pub fn load_config() -> Result<Config, Box<dyn Error>> {
    let path = config_path()?;
    if !path.exists() {
        return Err("Configuration file not found. Please run `asars config` to set it up.".into());
    }

    let config_data = fs::read_to_string(path)?;
    let config: Config = serde_json::from_str(&config_data)?;
    Ok(config)
}

pub fn get_token() -> Result<String, Box<dyn std::error::Error>> {
    let entry = Entry::new("asars", "token")?;

    match entry.get_password() {
        Ok(token) => Ok(token),
        Err(keyring::Error::NoEntry) => Err(Box::new(std::io::Error::new(
            std::io::ErrorKind::NotFound,
            "No token found. Please run `asars config` to set up your Asana token.",
        ))),
        Err(e) => Err(Box::new(e)),
    }
}

pub fn delete_token() -> Result<String, Box<dyn std::error::Error>> {
    let entry = Entry::new("asars", "token")?;
    entry.delete_credential()?;
    Ok("Token deleted.".to_string())
}
