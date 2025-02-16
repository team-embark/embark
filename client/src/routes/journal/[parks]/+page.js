export async function load({ params }) {
    const { parks } = params;

    // Mock database or API response (you can replace this with an actual fetch)
    const parksData = {
        "joshua-tree": {
            title: "Joshua Tree National Park",
            image: "JoshuaTreeirlpicture.png",
            modalImage: "joshuatreeafter3.jpg",
            logoImage: "JoshuaTree-2.png",
            facts: [
                { fact: "Named after the unique Joshua Tree.", icon: "🌵" },
                { fact: "Home to over 250 species of birds.", icon: "🦅" },
                { fact: "Designated as a National Park in 1994.", icon: "🎉" },
                { fact: "Over 8,000 years of human history, including Native American rock art.", icon: "🖼️" }
            ],
            modalText: `
                <h2>🔥 The Impact of Wildfires on Joshua Tree National Park 🔥</h2>
                <p><strong>The York Fire (2023)</strong> burned thousands of acres, devastating the fragile ecosystem.</p>
                <ul>
                    <li>🌵 Joshua trees are not fire-resistant, and large wildfires threaten their survival.</li>
                    <li>🌡 Climate change and human activity increase fire frequency and intensity.</li>
                    <li>🚫 Once burned, Joshua trees rarely grow back in affected areas.</li>
                    <li>🛑 Conservation efforts, responsible tourism, and fire safety awareness are crucial.</li>
                </ul>
                <p>Protecting our national parks starts with <strong>YOU</strong>.</p>
            `
        }
    };

    // Get the park data based on the slug
    const pd = parksData[parks] || null;

    if (!pd) {
        return { status: 404, error: new Error("Park not found") };
    }

    return {
        park: pd
    };
}
