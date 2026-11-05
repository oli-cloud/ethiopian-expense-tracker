# Project: Personal Expense Tracker – ETB Edition

Act as a Senior Full-Stack Software Engineer, UI/UX Designer, and Software Architect.

I want to develop a modern, responsive, user-friendly Personal Expense Tracker application designed primarily for users in Ethiopia.

The application should help individuals manage their daily income and expenses, categorize transactions, monitor budgets, and understand their monthly financial activities.

## 1. Project Objectives

Build a professional personal finance management application with:

- Daily income and expense tracking.
- Expense categorization.
- Monthly financial summaries.
- Interactive charts and analytics.
- Budget management.
- Ethiopian Birr (ETB) support.
- English and Amharic language support.
- Mobile-friendly and desktop-friendly interfaces.

## 2. Technology Stack

Use the following technologies:

### Frontend
- React.js
- JavaScript
- Tailwind CSS
- React Router
- Recharts for data visualization
- Lucide React for icons

### Backend (later development phase)
- Go programming language
- Gin framework
- RESTful API architecture

### Database (later development phase)
- PostgreSQL
- SQLite for initial development if necessary

### Initial Storage
- LocalStorage

Important: Start with a frontend-only application using LocalStorage. Design the architecture so that it can later integrate with a Go backend and PostgreSQL without rewriting the entire frontend.

## 3. Main Features

### A. Dashboard

Create a modern dashboard displaying:

- Total income.
- Total expenses.
- Available balance.
- Monthly savings.
- Recent transactions.
- Monthly income versus expense chart.
- Expense distribution by category.
- Spending trends.

### B. Transaction Management

Users should be able to:

- Add income transactions.
- Add expense transactions.
- Edit transactions.
- Delete transactions.
- Search transactions.
- Filter by date, category, and transaction type.
- View transaction history.

Transaction fields:

- Transaction title.
- Amount.
- Transaction type (Income or Expense).
- Category.
- Date.
- Payment method.
- Description.

### C. Expense Categories

Provide default categories:

- Food and groceries.
- Transportation.
- Housing and rent.
- Electricity and water.
- Internet and airtime.
- Healthcare.
- Education.
- Clothing.
- Entertainment.
- Family expenses.
- Other.

Allow users to create, edit, and delete custom categories.

### D. Income Management

Include income sources such as:

- Salary.
- Business income.
- Freelancing.
- Other income.

### E. Budget Management

Allow users to:

- Set monthly budgets.
- Set budgets for individual categories.
- Monitor budget utilization.
- View remaining budget.
- Receive warnings when approaching or exceeding budgets.

### F. Reports and Analytics

Provide:

- Daily reports.
- Weekly reports.
- Monthly reports.
- Annual reports.
- Category-based expense reports.
- Income versus expense comparisons.
- Monthly savings analysis.

Use Recharts to create attractive and interactive charts.

### G. Payment Methods

Support common Ethiopian payment methods:

- Cash.
- Telebirr.
- CBE Birr.
- Bank transfer.
- Other.

### H. Language Support

Implement English and Amharic language switching.

Use a clean translation structure so additional languages can be added easily.

English should be the default language.

### I. Settings

Include:

- Currency selection (ETB by default).
- Language preferences.
- Theme preferences (light/dark).
- Data export.
- Data backup and restore.
- Reset application data.

## 4. UI/UX Design Requirements

Design a professional financial dashboard inspired by modern fintech applications.

Design principles:

- Clean and minimal interface.
- Modern typography.
- Consistent spacing.
- Professional color palette.
- Responsive layout.
- Accessible forms and controls.
- Clear navigation.
- Attractive dashboard cards.
- Interactive charts.
- Useful empty states.
- Loading and error states where applicable.
- Confirmation before destructive actions.

Use a sidebar navigation on desktop and a mobile-friendly navigation system on smaller screens.

Avoid unnecessary animations, excessive colors, and complicated interfaces.

## 5. Application Pages

Create the following pages:

1. Dashboard
2. Transactions
3. Add Transaction
4. Categories
5. Budgets
6. Reports and Analytics
7. Settings
8. About

## 6. Architecture and Code Quality

Follow these development principles:

- Use reusable React components.
- Organize code into logical folders.
- Separate business logic from UI components.
- Use reusable custom hooks where appropriate.
- Implement centralized transaction calculations.
- Avoid duplicated code.
- Use consistent naming conventions.
- Validate all user inputs.
- Handle invalid amounts and dates.
- Prevent negative or invalid transaction amounts.
- Use proper error handling.
- Keep the code maintainable and scalable.

Suggested structure:

src/
  components/
  pages/
  layouts/
  hooks/
  services/
  utils/
  context/
  locales/
  assets/

## 7. Data Management

For the initial version:

- Store transactions in LocalStorage.
- Store categories in LocalStorage.
- Store budgets in LocalStorage.
- Persist user preferences.
- Ensure data remains available after refreshing the browser.
- Implement JSON backup and restore.

Do not store sensitive financial or authentication information insecurely.

## 8. Future Backend Integration

Prepare the frontend for future integration with:

- Go.
- Gin.
- PostgreSQL.
- REST APIs.
- JWT authentication.

Create a service layer so LocalStorage operations can later be replaced with API requests without significantly changing the UI components.

## 9. Development Instructions

Work incrementally.

Before writing code:

1. Analyze the requirements.
2. Propose the application architecture.
3. Define the database schema for future backend integration.
4. Identify the reusable components.
5. Propose the page navigation and user workflows.
6. Explain the implementation plan.

Then implement the application step by step.

Do not generate the entire application in one unnecessarily large response.

Start by building a complete, functional frontend MVP.

Do not introduce unnecessary dependencies.

Do not modify or remove existing project files without first inspecting the project structure.

Make sure every visible button and form performs its intended action.

## 10. Final Deliverable

I expect a working application that:

- Looks like a professional fintech product.
- Works on mobile, tablet, and desktop.
- Supports Ethiopian Birr.
- Supports English and Amharic.
- Tracks income and expenses.
- Provides meaningful financial analytics.
- Saves data persistently using LocalStorage.
- Has a clean architecture ready for Go backend integration.

Begin by presenting the architecture, folder structure, and implementation plan. Then start implementing the frontend MVP.