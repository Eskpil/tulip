abstract class Subcommand {
  Map<String, dynamic> toJson();
}

class Command {
  final Subcommand subcommand;
  final String subject;

  const Command({
    required this.subject,
    required this.subcommand,
  });

  Map<String, dynamic> toJson() => {
    'subject': subject,
    'subcommand': subcommand.toJson(),
  };
}