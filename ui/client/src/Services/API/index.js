const makeResult = (error, data) => ({
  error,
  data,
});

export default async (input, init) => {
  try {
    const response = await fetch(input, init);
    if (response.ok) {
      const data = await response.json();
      return makeResult(null, data);
    }

    return makeResult(`${response.status}`, null);
  } catch (e) {
    return makeResult(e, null);
  }
};
