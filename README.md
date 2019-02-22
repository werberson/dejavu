# Dejavu

This project is a small experiment to showcase the realtime capabilities of Prometheus.

It consists of a simple webpage that identifies you browser and device. 

The webpage communicates every second with the backend so that a bug can be installed via the `/api/bug/{platform}` API.

This bug makes the requests of the indicated `platform` to slow down to near `5s`.

