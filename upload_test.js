import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '3s', target: 30 },
    { duration: '10s', target: 70 },
    { duration: '2s', target: 10 },
  ],
};

// open() is GLOBAL in older k6
const video = open('./video.mp4', 'b');

export default function () {
  const res = http.post(
    'http://localhost:8080/upload',
    {
      file: http.file(video, 'video.mp4', 'video/mp4'),
    },
    { timeout: '120s' }
  );

  check(res, {
    'upload ok': (r) => r.status === 200 || r.status === 201 || r.status === 204,
  });

  sleep(1);
}
