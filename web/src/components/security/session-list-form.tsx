"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";

interface Session {
  id: number;
  device: string;
  location: string;
  last_login: string;
}

const fakeSessions: Session[] = [
  {
    id: 1,
    device: "Chrome on Windows",
    location: "New York, USA",
    last_login: "2025-02-20 15:00:00",
  },
  {
    id: 2,
    device: "Safari on iPhone",
    location: "London, UK",
    last_login: "2025-02-19 10:30:00",
  },
  {
    id: 3,
    device: "Firefox on Linux",
    location: "Berlin, Germany",
    last_login: "2025-02-18 08:20:00",
  },
];

export function SessionList() {
  const [sessions, setSessions] = useState<Session[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    setSessions(fakeSessions);
  }, []);
  async function revokeSession(id: number) {
    setSessions((prevSessions) =>
      prevSessions.filter((session) => session.id !== id)
    );
  }
  return (
    <div>
      {error && <p className="text-red-500 text-xs mt-1">{error}</p>}
      <ul className="divide-y">
        {sessions.map((session) => (
          <li
            key={session.id}
            className="py-2 flex justify-between items-center"
          >
            <div>
              <p>
                {session.device} - {session.location}
              </p>
              <p className="text-sm text-muted-foreground">
                Last login: {session.last_login}
              </p>
            </div>

            <Button onClick={() => revokeSession(session.id)}>Revoke</Button>
          </li>
        ))}
      </ul>
    </div>
  );
}
