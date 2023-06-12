export function splitArrayIntoGroups<T>(array: T[], n: number) {
  const groups = [];
  const size = Math.ceil(array.length / n);
  for (let i = 0; i < array.length; i += size) {
    groups.push(array.slice(i, i + size));
  }
  return groups;
}

export async function wait(ms: number) {
  return new Promise((resolve) => {
    setTimeout(() => resolve(null), ms);
  });
}
