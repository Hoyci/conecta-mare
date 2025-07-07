import Navbar from "@/components/layout/Navbar";
import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

const DashboardSkeleton = () => {
  return (
    <div className="min-h-screen flex flex-col bg-conecta-gray">
      <Navbar />
      <main className="flex-grow py-6">
        <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
          {/* Skeleton do Header do Perfil */}
          <Card className="border-none shadow-md">
            <CardContent className="p-0">
              <div className="flex flex-col py-4 md:flex-row md:items-end px-6 relative">
                <div className="w-full flex flex-col gap-4">
                  <div className="flex flex-row items-center gap-6">
                    <Skeleton className="w-24 h-24 rounded-full" />
                    <div className="flex-1 space-y-3">
                      <Skeleton className="h-7 w-48 rounded-md" />
                      <Skeleton className="h-5 w-64 rounded-md" />
                      <Skeleton className="h-5 w-72 rounded-md" />
                    </div>
                    <Skeleton className="h-11 w-32 rounded-lg" />
                  </div>

                  <div className="w-full flex gap-4 justify-between">
                    <Skeleton className="h-24 w-full rounded-lg" />
                    <Skeleton className="h-24 w-full rounded-lg" />
                    <Skeleton className="h-24 w-full rounded-lg" />
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Skeleton das Abas */}
          <div className="flex gap-2">
            <Skeleton className="h-10 w-28 rounded-md" />
            <Skeleton className="h-10 w-28 rounded-md" />
            <Skeleton className="h-10 w-28 rounded-md" />
          </div>

          {/* Skeleton do Conteúdo das Abas */}
          <div className="grid md:grid-cols-2 gap-6">
            <Card className="border-none shadow-md h-[400px]">
              <CardContent className="p-6">
                <Skeleton className="h-full w-full" />
              </CardContent>
            </Card>
            <div className="grid gap-6">
              <Card className="border-none shadow-md h-[188px]">
                <CardContent className="p-6">
                  <Skeleton className="h-full w-full" />
                </CardContent>
              </Card>
              <Card className="border-none shadow-md h-[188px]">
                <CardContent className="p-6">
                  <Skeleton className="h-full w-full" />
                </CardContent>
              </Card>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default DashboardSkeleton;

