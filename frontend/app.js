async function fetchData() {
    try {
        const response = await fetch('/api/status');
        const data = await response.json();
        document.getElementById('server-name').innerText = data.server_name;
        document.getElementById('server-time').innerText = data.time;
        const cacheHeader = response.headers.get('X-Cache-Status');
        document.getElementById('cache-status').innerText = cacheHeader ? cacheHeader : "Не налаштовано";
    } catch (error) {
        console.error("Помилка отримання даних:", error);
        document.getElementById('server-name').innerText = "Помилка зв'язку з API";
    }
}

fetchData();
