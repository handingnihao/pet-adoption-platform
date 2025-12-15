import { createBrowserRouter } from 'react-router-dom'
import { Layout } from '../components/layout/Layout'
import { AdminLayout } from '../components/layout/AdminLayout'
import { Home } from '../pages/Home'
import { Login } from '../pages/Login'
import { Register } from '../pages/Register'
import { Profile } from '../pages/Profile'
import { PetList } from '../pages/PetList'
import { PetDetail } from '../pages/PetDetail'
import { AdoptApply } from '../pages/AdoptApply'
import { MyApplications } from '../pages/MyApplications'
import { MyPets } from '../pages/MyPets'
import { CreatePet } from '../pages/CreatePet'
import { Community } from '../pages/Community'
import { PostDetail } from '../pages/PostDetail'
import { CreatePost } from '../pages/CreatePost'
import { Guide } from '../pages/Guide'
import { About } from '../pages/About'
import { AdminDashboard } from '../pages/admin/Dashboard'
import { UserManagement } from '../pages/admin/UserManagement'
import { PetManagement } from '../pages/admin/PetManagement'
import { AdoptionManagement } from '../pages/admin/AdoptionManagement'
import { OrganizationReview } from '../pages/admin/OrganizationReview'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { index: true, element: <Home /> },
      { path: 'login', element: <Login /> },
      { path: 'register', element: <Register /> },
      { path: 'profile', element: <Profile /> },
      { path: 'pets', element: <PetList /> },
      { path: 'pets/create', element: <CreatePet /> },
      { path: 'pets/:id', element: <PetDetail /> },
      { path: 'adopt/:id', element: <AdoptApply /> },
      { path: 'my-applications', element: <MyApplications /> },
      { path: 'my-pets', element: <MyPets /> },
      { path: 'community', element: <Community /> },
      { path: 'community/create', element: <CreatePost /> },
      { path: 'community/:id', element: <PostDetail /> },
      { path: 'guide', element: <Guide /> },
      { path: 'about', element: <About /> },
    ],
  },
  {
    path: '/admin',
    element: <AdminLayout />,
    children: [
      { index: true, element: <AdminDashboard /> },
      { path: 'users', element: <UserManagement /> },
      { path: 'pets', element: <PetManagement /> },
      { path: 'adoptions', element: <AdoptionManagement /> },
      { path: 'organizations', element: <OrganizationReview /> },
    ],
  },
])
