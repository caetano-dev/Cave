# Notetaking app

Work in progress.

## Running the project

Install the dependencies:

From the backend folder:
`go mod tidy`

From the frontend folder:
`npm install`

Run the backend with `DATABASE_NAME=database.db go run src/main.go` and upload the files with the `index.html` page from the backend folder.
Run the frontend with `npm run dev`.

To-do (for now):

- Fix the last test function for the database.
- Fix the sidebar that does not have a scroll.
- Let user create and delete files.
- Let user change file names.
- Create a search bar where the user can look for files based on the filename and tags.

---

### What are we building?

The goal of this project is to develop a local-first, FOSS and self-hosted notetaking app, where you can create and modify text files, images and documents. In the future, we want to let users access their files from anywhere, by allowing them to host the software on a dedicated machine, such as a Raspberry pi. 

### Why are we building it?

Notion is cool, but it requires internet access. Obsidian is great, but it is proprietary and you need to pay to sync your files. Evernote lets you create as many notes as you want, but it's not very privacy friendly. Other notetaking apps often add a limit of how many files you can create on a free plan. After spending so much time looking for the perfect solution to our problems, we have decided to try to create it ourselves.

### How are we building it?

The two main technologies we are using for the moment are React and Go. If you have programming knowledge and want to help, open a pull request!
