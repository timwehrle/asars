use crate::commands::config::load_config;
use crate::utils::format_date::format_date;
use reqwest::Client;
use serde_json::Value;
use url::Url;

#[derive(Debug)]
struct Task {
    name: String,
    due_on: Option<String>,
}

impl Task {
    fn from_json(task_json: &Value) -> Self {
        Self {
            name: task_json["name"]
                .as_str()
                .unwrap_or("Unnamed Task")
                .to_string(),
            due_on: task_json["due_on"]
                .as_str()
                .map(|s| format_date(s, "%b %d, %Y"))
                .or(Some("None".to_string())),
        }
    }
}

fn render_tasks(tasks: &[Task]) {
    println!("\nYour Tasks\n");
    tasks.iter().enumerate().for_each(|(i, task)| {
        println!(
            "{}. [{}] {}\n",
            i + 1,
            task.due_on.as_deref().unwrap_or("None"),
            task.name
        );
    });
}

pub async fn list_tasks(token: &str) -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new();
    let config = load_config().map_err(|err| {
        eprintln!("Failed to load config: {}", err);
        err
    })?;

    let base_url = "https://app.asana.com/api/1.0/tasks";
    let url = Url::parse_with_params(
        base_url,
        &[
            ("assignee", "me"),
            ("workspace", &config.workspace),
            ("completed_since", "now"),
            ("opt_fields", "name,completed,due_on"),
        ],
    )?;

    let response = client.get(url).bearer_auth(token).send().await?;
    if !response.status().is_success() {
        eprintln!(
            "Error: Failed to fetch tasks. Status: {}",
            response.status()
        );
        return Err("Failed to fetch tasks".into());
    }

    let json: Value = response.json().await?;
    if let Some(task_array) = json["data"].as_array() {
        let tasks: Vec<Task> = task_array.iter().map(Task::from_json).collect();
        render_tasks(&tasks);
    } else {
        println!("No tasks found.");
    }
    Ok(())
}
