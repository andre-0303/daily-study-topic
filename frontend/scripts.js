const title = document.getElementById("title");
const difficulty = document.getElementById("difficulty");
const description = document.getElementById("description");
const catalog = document.getElementById("catalog");
const button = document.getElementById("generate-btn");
const loadingTip = document.getElementById("loading-tip");
const card = document.querySelector(".card");

const API_URL =
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1"
    ? "http://localhost:8080/api/topic"
    : "https://daily-study-topic.onrender.com/api/topic";

const tips = [
  "Go compila para um binário único e estático — sem runtime externo, sem dependências no servidor.",
  "Goroutines custam ~2KB de memória. Threads do sistema operacional custam ~1MB cada.",
  "Em Go, se um identificador começa com letra maiúscula, ele é exportado (público). Minúscula = privado ao pacote.",
  "Go não tem classes nem herança. Composição via structs e interfaces cobre tudo.",
  "Channels em Go são filas thread-safe para comunicação entre goroutines — preferidos a mutexes para dados compartilhados.",
  "O garbage collector do Go roda concorrentemente com o programa, sem parar a execução.",
  "go:embed permite incluir arquivos estáticos (SQL, HTML, configs) direto no binário em tempo de compilação.",
  "defer executa a função no retorno do escopo — útil para fechar arquivos, conexões e locks sem esquecer.",
  "Go trata erros como valores: funções retornam (resultado, error). Não há exceções.",
  "O compilador do Go rejeita imports e variáveis declarados mas não usados — código limpo por padrão.",
  "interfaces em Go são implícitas — um tipo satisfaz uma interface só de ter os métodos, sem declarar.",
  "go test roda benchmarks com -bench=. e mostra ns/op — benchmarking embutido na toolchain.",
  "O scheduler do Go mapeia N goroutines para M threads do OS — modelo M:N de concorrência.",
  "make() e new() são distintos: new aloca e zera, make inicializa slices, maps e channels.",
];

let tipInterval = null;
let tipIndex = 0;

function startTips() {
  tipIndex = Math.floor(Math.random() * tips.length);
  loadingTip.textContent = tips[tipIndex];
  card.classList.add("is-loading");

  tipInterval = setInterval(() => {
    tipIndex = (tipIndex + 1) % tips.length;
    loadingTip.textContent = tips[tipIndex];
    loadingTip.style.animation = "none";
    loadingTip.offsetHeight;
    loadingTip.style.animation = "";
  }, 3500);
}

function stopTips() {
  clearInterval(tipInterval);
  tipInterval = null;
  card.classList.remove("is-loading");
}

async function fetchTopic() {
  button.disabled = true;
  card.classList.add("is-updating");
  startTips();

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
  } finally {
    stopTips();
    card.classList.remove("is-updating");
    button.disabled = false;
  }
}

button.addEventListener("click", fetchTopic);

fetchTopic();
