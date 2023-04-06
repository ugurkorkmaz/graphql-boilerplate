import "./style/reset.css";
import App from "./App.svelte";

const gqlgen = Array.from(document.getElementsByTagName("gqlgen"));

gqlgen.forEach((el: Element) => {
  // data is a stringified JSON object or never
  const data = el.getAttribute("data-props");

  // props is a JSON object or never
  const props = data ? JSON.parse(data) : {};

  const app = new App({
    target: el,
    props,
  });

  // app is a Svelte component or never
  if (app) {
    // app.$on is a function or never
    app.$on("destroy", () => {
      app.$destroy();
    });
  }

  return app;
});

export default App;
