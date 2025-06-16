
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Notification } from "@/types/metrics";
import { Bell, Award, Lightbulb } from "lucide-react";

interface NotificationsSectionProps {
  notifications: Notification[];
}

export function NotificationsSection({ notifications }: NotificationsSectionProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex justify-between items-center">
          <CardTitle className="text-lg">Notificações</CardTitle>
          <span className="inline-flex items-center justify-center w-6 h-6 bg-conecta-blue-light text-white text-xs font-bold rounded-full">
            {notifications.filter(n => !n.read).length}
          </span>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {notifications.map((notification) => (
            <div
              key={notification.id}
              className={`flex p-3 rounded-lg ${
                notification.read ? "bg-transparent" : "bg-blue-50"
              }`}
            >
              <div className="mr-4 mt-0.5">
                {notification.type === "milestone" && (
                  <Award className="h-6 w-6 text-conecta-blue" />
                )}
                {notification.type === "review" && (
                  <Bell className="h-6 w-6 text-yellow-500" />
                )}
                {notification.type === "tip" && (
                  <Lightbulb className="h-6 w-6 text-conecta-green" />
                )}
              </div>
              <div className="space-y-1">
                <p className="font-medium">{notification.title}</p>
                <p className="text-sm text-muted-foreground">
                  {notification.description}
                </p>
                <p className="text-xs text-muted-foreground">
                  {new Date(notification.date).toLocaleDateString("pt-BR")}
                </p>
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
