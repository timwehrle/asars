mod asana_api;
mod cli;
mod commands;
mod utils;

use crate::commands::config::{configure, get_token};
use crate::commands::tasks::list_tasks;
use clap::Parser;
use cli::{Commands, CLI};
use commands::config::handle_config_command;
use commands::projects;

#[tokio::main]
async fn main() {
    let cli = CLI::parse();

    match &cli.command {
        Commands::Config { delete } => {
            if *delete {
                handle_config_command(*delete)
                    .await
                    .expect("Failed to delete token.");
                return;
            }
            configure().await.expect("Failed to configure Asana CLI.");
        }

        Commands::Projects { list, get } => {
            if *list {
                let token = get_token().expect("Failed to get token: ");
                projects::list_projects(&token)
                    .await
                    .expect("Failed to list projects.");
            }

            if let Some(project_id) = get {
                println!("Getting project details for id: {}", project_id);
            }
        }

        Commands::Tasks => {
            let token = get_token().expect("Failed to get token: ");
            list_tasks(&token).await.expect("Failed to list tasks.");
        }
    }
}
