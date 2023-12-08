
const FILENAME = "example.txt"


async function main() {
  const file = Bun.file(FILENAME)
  const content = await file.text()
  console.log(content)
}


main()
