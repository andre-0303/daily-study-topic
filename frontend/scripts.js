const title = document.getElementById("title");
const difficulty = document.getElementById("difficulty");
const description = document.getElementById("description");
const catalog = document.getElementById("catalog");
const button = document.getElementById("generate-btn");

const API_URL =
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1"
    ? "http://localhost:8080/api/topic"
    : "https://daily-study-topic.onrender.com/api/topic";

async function fetchTopic() {
  try {
    const response = await fetch(API_URL);

    if (!response.ok) {
      throw new Error("Erro ao buscar tópico");
    }

    const topic = await response.json();

    title.textContent = topic.title;
    difficulty.textContent = topic.difficulty;
    description.textContent = topic.description;
    catalog.textContent = "Nº " + String(topic.id).padStart(3, "0");
  } catch (error) {
    title.textContent = "Erro ao carregar";
    difficulty.textContent = "";
    description.textContent = "Não foi possível carregar os dados.";
    catalog.textContent = "Nº 000";

    console.error(error);
  }
}

button.addEventListener("click", fetchTopic);

fetchTopic();
