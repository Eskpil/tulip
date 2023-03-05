
import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';

import '../models/entity.dart';

@immutable
abstract class EntitiesState extends Equatable {
  const EntitiesState();
}

class EntitiesLoading extends EntitiesState {
  @override
  List<Object> get props => [];
}

class EntitiesLoaded extends EntitiesState {
  final List<Entity> entities;

  const EntitiesLoaded({ required this.entities });

  @override
  List<Object> get props => [];
}

class EntitiesError extends EntitiesState {
  @override
  List<Object> get props => [];
}