import React, { Suspense, lazy } from "react";
const Dashboard = lazy(() => import("./views/Dashboard"));
const Login = lazy(() => import("./views/Login"));
const ResponseSelection = lazy(() => import("./views/ResponseSelection"));
const Score = lazy(() => import("./views/Score"));
const Header = lazy(() => import("./views/Header"));

function App() {
  const location = useLocation();
  const isLoggedIn = localStorage.getItem("userId");

  return (
    <>
      <FetchUserDetails />
      {isLoggedIn && location.pathname !== "/" && (
        <Suspense fallback={<div>Loading...</div>}>
          <Header />
        </Suspense>
      )}
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route index path="/" element={<Login />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/responses/:cardId" element={<ResponseSelection />} />
          <Route path="/score" element={<Score />} />
        </Routes>
      </Suspense>
    </>
  );
}
export default App