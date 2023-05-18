// @ts-check

/** @type {HTMLFormElement | null} */
const form = document.querySelector("#create-form");

if (form !== null) {
  form.addEventListener("submit", async (ev) => {
    ev.preventDefault();

    const data = new FormData(form);

    const res = await fetch("/api/snippets", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: data.get("name"),
        content: data.get("content"),
      }),
    });

    if (res.ok) {
      const svrData = await res.json();

      // TODO: Create frontend for viewing snippets
      window.location.assign(`/api/snippets/${svrData["id"]}`);
    } else {
      console.error(`Failed creating snippet: ${res.body}`);
    }
  });
}
