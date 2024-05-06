import { tsParticles } from "https://cdn.jsdelivr.net/npm/@tsparticles/engine@3.0.3/+esm";
import { loadAll } from "https://cdn.jsdelivr.net/npm/@tsparticles/all@3.0.3/+esm";

(async ParticlesBackground => {
    await loadAll(tsParticles);

    await tsParticles.addPreset("lightdark", {
        fullScreen: {
            enable: true
        },
        particles: {
            links: {
                enable: true
            },
            move: {
                enable: true
            },
            number: {
                value: 50
            },
            opacity: {
                value: { min: 0.3, max: 1 }
            },
            shape: {
                type: ["circle", "square", "triangle", "polygon"],
                options: {
                    polygon: [
                        {
                            sides: 5
                        },
                        {
                            sides: 6
                        },
                        {
                            sides: 8
                        }
                    ]
                }
            },
            size: {
                value: { min: 1, max: 3 }
            }
        }
    });

    await tsParticles.load({
        id: "light",
        options: {
            preset: "lightdark",
            particles: {
                color: {
                    value: "#3b82f6"
                },
                links: {
                    color: "#3b82f6"
                }
            }
        }
    });
})();
