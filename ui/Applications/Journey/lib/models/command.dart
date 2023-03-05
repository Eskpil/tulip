abstract class Subcommand {
  Map<String, dynamic> toJson();
}

class Command {
  final Subcommand subcommand;

  const Command({
    required this.subcommand,
  });

  Map<String, dynamic> toJson() => {
    'subcommand': subcommand.toJson(),
  };
}