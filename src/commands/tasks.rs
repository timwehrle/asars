use crate::commands::config::load_config;
use crate::utils::format_date::format_date;
use reqwest::Client;
use serde_json::Value;

#[derive(Debug)]
struct Task {
    name: String,
    due_on: Option<String>,
}

impl Task {
    fn from_json(task_json: &Value) -> Self {
        Task {
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
    println!("\n--- Your Tasks ---");
    for (i, task) in tasks.iter().enumerate() {
        println!(
            "\n{}. [{}] {}",
            i + 1,
            task.due_on.as_deref().unwrap_or("None"),
            task.name
        );
    }
    println!("\n------------------\n");
}

pub async fn list_tasks(token: &str) -> Result<(), Box<dyn std::error::Error>> {
    let client = Client::new();
    let config = load_config().expect("Could not load config");

    let url = format!(
        "https://app.asana.com/api/1.0/tasks?assignee=me&workspace={}&completed_since=now&opt_fields=name,completed,due_on",
        config.workspace
    );
    let response = client.get(&url).bearer_auth(token).send().await?;

    let json: Value = response.json().await?;

    if let Some(task_array) = json["data"].as_array() {
        let tasks: Vec<Task> = task_array.iter().map(Task::from_json).collect();

        render_tasks(&tasks);
    } else {
        println!("No tasks found.");
    }
    Ok(())
}
