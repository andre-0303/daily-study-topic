const title = document.getElementById("title");
const difficulty = document.getElementById("difficulty");
const description = document.getElementById("description");
const catalog = document.getElementById("catalog");
const card = document.querySelector(".card");
const button = document.getElementById("generate-btn");

async function fetchTopic() {
  try {
    card.classList.remove("is-updating");
    void card.offsetWidth;

    const response = await fetch("http://localhost:8080/api/topic");

    const topic = await response.json();

    title.textContent = topic.title;
    difficulty.textContent = topic.difficulty;
    description.textContent = topic.description;
    catalog.textContent = "Nº " + String(topic.id).padStart(3, "0");

    card.classList.add("is-updating");
  } catch (error) {
    title.textContent = "Erro ao carregar";
    console.error(error);
  }
}

button.addEventListener("click", fetchTopic);

fetchTopic();
