
import { Avatar } from "@/components/ui/avatar";
import { AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Card } from "@/components/ui/card";

interface ProfileHeaderProps {
  name: string;
  profession: string;
  avatar: string;
  description: string;
  location: string;
  coverPhoto: string;
}

export function ProfileHeader({
  name,
  profession,
  avatar,
  description,
  location,
  coverPhoto,
}: ProfileHeaderProps) {
  return (
    <Card className="overflow-hidden border-none shadow-md">
      <div 
        className="h-40 w-full bg-cover bg-center"
        style={{ backgroundImage: `url(${coverPhoto})` }}
      >
        <div className="h-full w-full bg-gradient-to-t from-black/60 to-transparent" />
      </div>
      <div className="relative -mt-16 px-6 pb-6">
        <Avatar className="h-32 w-32 border-4 border-white shadow-lg">
          <AvatarImage src={avatar} alt={name} />
          <AvatarFallback className="bg-conecta-blue text-white text-2xl">
            {name.charAt(0)}
          </AvatarFallback>
        </Avatar>
        <div className="mt-4 flex flex-col gap-1 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="text-2xl font-bold text-gray-900">{name}</h1>
            <p className="text-lg font-medium text-conecta-blue">{profession}</p>
            <p className="mt-1 text-gray-500">{location}</p>
          </div>
          <div className="mt-3 flex gap-2 sm:mt-0">
            <button className="btn-primary">Mensagem</button>
            <button className="bg-gray-200 text-gray-800 px-4 py-2 rounded-md font-medium hover:bg-gray-300 transition-colors">
              Compartilhar
            </button>
          </div>
        </div>
        <p className="mt-4 text-gray-700">{description}</p>
      </div>
    </Card>
  );
}
