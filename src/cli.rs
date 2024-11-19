use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(name = "asars")]
#[command(about = "A CLI for interacting with Asana.", long_about = None)]
pub struct CLI {
    #[command(subcommand)]
    pub command: Commands,
}

#[derive(Subcommand)]
pub enum Commands {
    Config {
        #[arg(short, long)]
        delete: bool,
    },

    Projects {
        #[arg(short, long)]
        list: bool,

        #[arg(short, long)]
        get: Option<String>,
    },

    #[command(alias = "ts")]
    Tasks,
}
