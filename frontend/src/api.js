// const req = (url) => fetch(`http://localhost:8080${url}`, { mode: 'no-cors' }).then((it) => it.json());

const req = (url) => {
	console.log({ url });
  switch (url) {
    case '/runs':
      console.log('listing runs');
      return Promise.resolve([{ id: '62dff5cd-8671-445a-9eec-53d236673368', name: 'Sep 29 e2e tests', status: 'failed' }]);
    case '/runs/62dff5cd-8671-445a-9eec-53d236673368':
      console.log('getting run');
      return Promise.resolve({ id: '62dff5cd-8671-445a-9eec-53d236673368', name: 'Sep 29 e2e tests', status: 'failed' });
    case '/runs/62dff5cd-8671-445a-9eec-53d236673368/suites':
      console.log('listing suites');
      return Promise.resolve([
        { id: '53df4ae0-867c-4eeb-832e-918ab277fbdf', name: 'Project page feature', status: 'failed' },
      ]);
    case '/runs/62dff5cd-8671-445a-9eec-53d236673368/suites/53df4ae0-867c-4eeb-832e-918ab277fbdf':
      console.log('getting suite');
      return Promise.resolve({
        id: '53df4ae0-867c-4eeb-832e-918ab277fbdf',
        name: 'Project page feature',
        status: 'failed',
      });
  }
};

export const api = {
  runs: {
    all: () => req('/runs'),
    one: (id) => req(`/runs/${id}`),
  },
  suites: {
    all: (runId) => req(`/runs/${runId}/suites`),
    one: (runId, id) => req(`/runs/${runId}/suites/${id}`),
  },
};
